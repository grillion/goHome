angular.module("homebase-api.auth", []).factory(
  'HomebaseAuth', [ '$q', '$log', 'HomebaseApi', function($q, $log, HomebaseApi) {
    'use strict';

    var HomebaseAuth = {};

    /**
     * Get current user
     */
    HomebaseAuth.me = function() {
      return HomebaseApi.get("auth/");
    };

    /**
     * Auth login via REST API
     * @param username
     * @param password
     */
    HomebaseAuth.login = function (username, password) {
      return HomebaseApi.post("auth/login", null, {
        username: username,
        password: password
      });
    };

    /**
     * Logout, VIA REST DELETE on /api/auth
     */
    HomebaseAuth.logout = function () {
      return HomebaseApi.delete("auth/");
    };

    return HomebaseAuth;
  }]
);