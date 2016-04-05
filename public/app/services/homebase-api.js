angular.module("homebase-api", [])
  .factory('HomebaseApi', [ '$http', '$q', function($http, $q) {
    'use strict';

    var HomebaseApi = {};

    HomebaseApi.query = function (method, url, params, dataObj) {

      var request = {
        method: method,
        url: "/api/" + url
      };

      if (request.method != "GET") request.headers = {"Content-Type": "application/json"};
      if (!!params) request.params = params;
      if (!!dataObj) request.data = dataObj;

      return $http(request).then(function(response){
        if(response.data.errorMessage) return $q.reject(response.data.errorMessage);
        return response.data;
      })
      .catch(function(err){
        return $q.reject(err.data.errorMessage || "Unknown server error")
      })
    };

    /**
     * HTTP POST to API
     * @param apiUri
     * @param params
     * @param dataObj
     */
    HomebaseApi.post = function (apiUri, params, dataObj) {
      return HomebaseApi.query("POST", apiUri, params, dataObj);
    };

    /**
     * HTTP GET to API
     * @param apiUri
     * @param params
     */
    HomebaseApi.get = function (apiUri, params) {
      return HomebaseApi.query("GET", apiUri, params);
    };

    /**
     * HTTP POST to API
     * @param apiUri
     * @param params
     * @param dataObj
     */
    HomebaseApi.put = function (apiUri, params, dataObj) {
      return HomebaseApi.query("PUT", apiUri, params, dataObj);
    };
    /**
     * HTTP POST to API
     * @param apiUri
     * @param params
     * @param dataObj
     */
    HomebaseApi.delete = function (apiUri, params, dataObj) {
      return HomebaseApi.query("DELETE", apiUri, params, dataObj);
    };

    return HomebaseApi;

  }]);