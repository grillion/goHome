/**
 * @typedef {Object} UserDef
 * @property {ObjectID} _id
 * @property {String} name
 * @property {String} username
 * @property {String} email
 * @property {String} password
 * @property {String} created
 * @property {String} updated
 */

angular.module("homebase-api.users", []).factory(
  'HomebaseUsers', [ '$q', '$log', 'HomebaseApi', function($q, $log, HomebaseApi) {
    'use strict';

    var HomebaseUsers = {};

    /**
     * Create a new user entry
     * @param {UserDef} userDef User definition
     */
    HomebaseUsers.create = function (userDef) {
      return HomebaseApi.post("users/", null, userDef);
    };

    /**
     * Get all users
     */
    HomebaseUsers.getAll = function() {
      return HomebaseApi.get("users/");
    };

    /**
     * Get a single user by ID
     * @param id
     */
    HomebaseUsers.get = function(id) {
      return HomebaseApi.get("users/" + id);
    };

    /**
     * Update a user
     * @param {UserDef} userDef
     */
    HomebaseUsers.update = function(userDef){
      return HomebaseApi.put("users/", null, userDef);
    };

    /**
     * Remove a user entry
     * @param id
     */
    HomebaseUsers.remove = function(id){
      return HomebaseApi.delete("users/" + id);
    };

    return HomebaseUsers;
  }]
);