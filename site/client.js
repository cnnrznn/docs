var Client = {
    PushQ: [],
    Inflight: "",
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
                this.PushQ.push({type:1, pos:start+i, c:' '})
            }
        } else if ("insert" in op) {
            for (i=0; i<op.insert.length; i++) {
                this.PushQ.push({type:0, pos:start+i, c:op.insert[i]})
            }
        }
    }

    console.log(this.PushQ)
};
