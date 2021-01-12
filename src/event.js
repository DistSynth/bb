
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
