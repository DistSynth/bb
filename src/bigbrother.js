
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
		
		// create iframe container if necessary
		if ( ! iframe_container ) {
		
			// create post frame container	
			var div = document.createElement( 'div' );
			div.setAttribute( 'id', container_id );
			document.body.appendChild( div );
			iframe_container = document.getElementById( container_id );
		}		
		
		// create iframe and post data once its fully loaded.
		this.generateHiddenIframe( iframe_container, data );
	},
}
