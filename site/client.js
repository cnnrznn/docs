const pb = require('./editor_pb.js');
const editor = require('./editor_grpc_web_pb.js');

var Client = {
    Id: -1,
    PushQ: [],
    Inflight: "",
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
                pbOp.setSender(this.Id);
                pbOp.setVersion(this.Version);
                pbOp.setType(1);
                pbOp.setPos(start);
                this.PushQ.push(pbOp);
            }
        } else if ("insert" in op) {
            for (i=0; i<op.insert.length; i++) {
                var pbOp = new pb.Op();
                pbOp.setSender(this.Id);
                pbOp.setVersion(this.Version);
                pbOp.setType(0);
                pbOp.setPos(start + i);
                pbOp.setChar(op.insert[i]);
                this.PushQ.push(pbOp);
            }
        }
    }

    console.log(this.PushQ)
};

Client.Tick = function() {
};

Client.ec = new editor.EditorClient('http://localhost:8080');
//console.log(Client.ec);
Client.ec.join(new pb.JoinRequest(), {}, function(err, resp) {
    console.log(err, resp);
    Client.id = resp.getId();
});

Client.ec.state(new pb.Nil, {}, function(err, resp) {
    version = resp.getVersion();
    buffer = String.fromCharCode.apply(String, resp.getBuffer());

    console.log(version, buffer);
});

window.Client = Client
