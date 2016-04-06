angular.module("homebase-api.devices.types", []).factory(
  'HomebaseDeviceTypes', function() {
    'use strict';


    var HomebaseDeviceTypes = {};

    /**
     * List of device types supported by device key
     */
    HomebaseDeviceTypes.listSupported = function() {
      return "mpower"
    };

    return HomebaseDeviceTypes;
  }
);