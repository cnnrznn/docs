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
    console.log(err, resp.getId());
    Client.Id = resp.getId();
});

Client.ec.state(new pb.Nil, {}, function(err, resp) {
    version = resp.getVersion();
    buffer = resp.getBuffer();
    console.log(version, buffer);

    Client.Version = version
    if (buffer.length > 0) {
        decoder = new TextDecoder('utf-8');
        quill.setText(decoder.decode(buffer));
    }
});

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
};

setTimeout(function tick() {
    Client.Tick();
    setTimeout(tick, 500);
}, 0);

window.Client = Client
