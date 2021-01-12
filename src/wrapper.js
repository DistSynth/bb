
(function () {
    console.log("From fs")
    // execute commands global owa_cmds command queue
    if (typeof bbCommands === 'undefined') {
        var q = new BB.commandQueue();
    } else {
        //if (OWA.util.is_array(owa_cmds)) {
        var q = new BB.commandQueue();
        q.loadCmds(bbCommands);
        // }
    }

    window['bbCommands'] = q;
    window['bbCommands'].process();
})();
