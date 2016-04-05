angular.module("homebase-api.settings", []).factory(
  'HomebaseSettings', [ '$q', '$log', 'HomebaseApi', function($q, $log, HomebaseApi) {
    'use strict';

    var HomebaseSettings = {};

    /**
     * Create or update a setting
     * @param key
     * @param value
     */
    HomebaseSettings.put = function(key, value) {
      return HomebaseApi.put("settings/", null, {
        key: key,
        value: value
      });
    };

    /**
     * Get all settings
     */
    HomebaseSettings.getAll = function() {
      return HomebaseApi.get("settings/");
    };

    /**
     * Get a single setting by key
     * @param key
     */
    HomebaseSettings.get = function(key) {
      return HomebaseApi.get("settings/" + encodeURIComponent(key));
    };

    /**
     * Remove a device entry
     * @param key
     */
    HomebaseSettings.delete = function(key){
      return HomebaseApi.delete("settings/" + encodeURIComponent(key));
    };


    return HomebaseSettings;
  }]
);