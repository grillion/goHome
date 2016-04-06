/**
 * @typedef {Object} DeviceDef
 * @property {ObjectID} _id
 * @property {String} name
 * @property {String} type
 * @property {String} created
 * @property {String} updated
 */

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
     * Update a Device
     * @param {DeviceDef} userDef
     */
    HomebaseDevices.update = function(deviceDef){
      return HomebaseApi.put("devices/", null, deviceDef);
    };

    /**
     * Remove a user entry
     * @param id
     */
    HomebaseDevices.remove = function(id){
      return HomebaseApi.delete("devices/" + id);
    };

    return HomebaseDevices;
  }]
);