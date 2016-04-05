angular.module("homebase-api.devices", []).factory(
  'HomebaseDevices', [ '$q', '$log', 'HomebaseApi', function($q, $log, HomebaseApi) {
    'use strict';

    var HomebaseDevices = {};

    /**
     * Create a new device entry
     * @param deviceDef
     */
    HomebaseDevices.create = function (deviceDef) {
      return HomebaseApi.post("devices/", null, deviceDef);
    };

    /**
     * Get all devices
     */
    HomebaseDevices.getAll = function() {
      return HomebaseApi.get("devices/");
    };

    /**
     * Get a single device by ID
     * @param id
     */
    HomebaseDevices.get = function(id) {
      return HomebaseApi.get("devices/" + id);
    };

    /**
     * Remove a device entry
     * @param id
     */
    HomebaseDevices.delete = function(id){
      return HomebaseApi.delete("devices/" + id);
    };

    return HomebaseDevices;
  }]
);