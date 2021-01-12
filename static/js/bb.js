
var BB = {
    items: {}
}

BB.commandQueue = function () {
    console.log('Command Queue object created');
    var icmds = [];
    var isPaused = false;
}

BB.commandQueue.prototype = {
    loadCmds: function (cmds) {
        this.icmds = cmds;
    },
    process: function () {

        var that = this;
        var callback = function () {
            // when the handler says it's finished (i.e. runs the callback)
            // We check for more tasks in the queue and if there are any we run again
            if (that.icmds.length > 0) {
                that.process();
            }
        }
        this.push(this.icmds.shift(), callback);
    },

    push: function (cmd, callback) {
        var args = Array.prototype.slice.call(cmd, 1);
        var obj_name = 'BigTracker';
        var method = cmd[0];

        if (typeof window[obj_name] == "undefined") {
            window[obj_name] = new BB.tracker({ globalObjectName: obj_name });
        }
        window[obj_name][method].apply(window[obj_name], args);

        if (callback && (typeof callback == 'function')) {
            callback();
        }

    },
}

BB.tracker = function (options) {
    this.options = { logPage: true, }
    this.setEndpoint(window.bbUrl);
    if (options) {
        for (var opt in options) {
            this.options[opt] = options[opt];
        }
    }
}

BB.tracker.prototype = {
    siteId: '',
    setSiteId: function (siteId) {
        this.siteId = siteId;
    },

    trackPageView: function (url) {
        var event = new BB.event;
        //if (url) {
        event.set('page_url', this.getCurrentUrl());
        //}
        event.setEventType("bb.page_view");
        console.log(event);
    },

    getCurrentUrl: function () {
        return document.URL
    },

    getTrackerEndpoint: function () {
        var url = this.getEndpoint();
        return url + 'tracker';
    },

    getEndpoint: function () {
        return this.getOption('baseUrl');
    },

    setEndpoint: function (endpoint) {
        this.setOption('baseUrl', endpoint);
    },

    getOption: function (name) {
        if (this.options.hasOwnProperty(name)) {
            return this.options[name];
        }
    },

    setOption: function (name, value) {
        this.options[name] = value;
    },

    cdPost : function ( data ) {
		
		var container_id = "bbContainer";
		
		var iframe_container = document.getElementById( container_id );
		
		if ( ! iframe_container ) {
			var div = document.createElement( 'div' );
			div.setAttribute( 'id', container_id );
			document.body.appendChild( div );
			iframe_container = document.getElementById( container_id );
        }
        
		this.generateHiddenIframe( iframe_container, data );
	},
}
;
BB.event = function () {

    this.properties = {};
    this.id = '';
    this.siteId = '';
    this.set('ts', '');
}

BB.event.prototype = {
    get: function (name) {
        if (this.properties.hasOwnProperty(name)) {
            return this.properties[name];
        }
    },

    set: function (name, value) {
        this.properties[name] = value;
    },

    setEventType: function (eventType) {
        this.set("eventType", eventType);
    },

    getProperties: function () {
        return this.properties;
    },

    apply: function (properties) {
        for (param in properties) {
            if (properties.hasOwnProperty(param)) {
                this.set(param, properties[param]);
            }
        }
    },

    isSet: function (name) {
        if (this.properties.hasOwnProperty(name)) {
            return true;
        }
    }
}
;
(function () {
    console.log("From fs")
    if (typeof bbCommands === 'undefined') {
        var q = new BB.commandQueue();
    } else {
        var q = new BB.commandQueue();
        q.loadCmds(bbCommands);
    }

    window['bbCommands'] = q;
    window['bbCommands'].process();
})();
