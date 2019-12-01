const pb = require('./editor_pb.js');
const editor = require('./editor_grpc_web_pb.js');

var Client = {
    Id: -1,
    PushQ: [],
    Inflight: "",
    Version: 0,
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
                pbOp.setPos(start + i);
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

window.Client = Client
