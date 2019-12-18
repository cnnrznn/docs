const pb = require('./editor_pb.js');
const editor = require('./editor_grpc_web_pb.js');

var Client = {
    Id: -1,
    PushQ: [],
    Inflight: null,
    Version: 0,

    ec: null,
};

Client.Push = function(ops) {
    start = 0 // default beginning
    type = 0  // default insert

    for (i in ops) {
        op = ops[i]

        if ("retain" in op) {
            start = op.retain
        } else if ("delete" in op) {
            type = 1
            for (i=0; i<op.delete; i++) {
                var pbOp = new pb.Op();
                pbOp.setType(1);
                pbOp.setPos(start);
                this.PushQ.push(pbOp);
            }
        } else if ("insert" in op) {
            for (i=0; i<op.insert.length; i++) {
                var pbOp = new pb.Op();
                pbOp.setType(0);
                pbOp.setPos(start + i);
                pbOp.setChar(new TextEncoder().encode(op.insert[i]));
                this.PushQ.push(pbOp);
            }
        }
    }

    console.log(this.PushQ)
};

Client.ec = new editor.EditorClient('http://localhost:8080');

Client.ec.join(new pb.JoinRequest(), {}, function(err, resp) {
    console.log("Joined Server:", err, resp.getId());
    Client.Id = resp.getId();
});

Client.ec.state(new pb.Nil, {}, function(err, resp) {
    version = resp.getVersion();
    buffer = resp.getBuffer();
    console.log("DocState:", version, buffer);

    Client.Version = version
    if (buffer.length > 0) {
        quill.setText(new TextDecoder().decode(buffer));
    }
});

var transform = function(op, other) {
    if (op.getType() == 0 && other.type == 0) {
        if (other.pos <= op.getPos()) {
            op.setPos(op.getPos() + 1)
        }
    } else if (op.getType() == 0 && other.type == 1) {
        if (other.pos < op.getPos()) {
            op.setPos(op.getPos() - 1)
        }
    } else if (op.getType() == 1 && other.type == 0) {
        if (other.pos <= op.getPos()) {
            op.setPos(op.getPos() + 1)
        }
    } else if (op.getType() == 1 && other.type == 1) {
        if (other.pos < op.getPos()) {
            op.setPos(op.getPos() - 1)
        } else {
            op.setType(2)
        }
    }
};

Client.Tick = function() {
    if (this.Inflight == null && this.PushQ.length > 0) {
        this.Inflight = this.PushQ.shift();
        this.Inflight.setVersion(this.Version);
        this.Inflight.setSender(this.Id);
        console.log(this.Inflight.toObject());

        this.ec.send(Client.Inflight, {}, function(err, resp) {
            console.log("Did server receive inflight?:", err, resp);
        });
    }

    var version = new pb.Version();
    version.setVersion(this.Version);

    var stream = this.ec.recv(version, {});
    stream.on('data', function(resp) {
        console.log(resp.toObject());
        if (resp.getSender() == Client.Id) {
            Client.Inflight = null;
        } else {
            for (i = 0; i<Client.PushQ.length; i++) {
                copyOp = resp.toObject()
                transform(resp, Client.PushQ[i].toObject())
                transform(Client.PushQ[i], copyOp)
            }
        }

        // apply resp to quill

    });
    stream.on('status', function(status) {
    });
    stream.on('end', function(end) {
    });
};

setTimeout(function tick() {
    Client.Tick();
    setTimeout(tick, 500);
}, 0);

window.Client = Client
