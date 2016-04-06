var module = angular.module("app.users", []);

module.directive('userForm', function(){
  return {
    restrict: "E",
    replace: true,
    templateUrl: "app/views/user/userForm.html",
    scope: {
      user: "=",
      onSubmit: "="
    }
  };
});

/**
 * List Users
 */
module.controller("UserListCtrl", ['$rootScope', '$scope', '$state', 'HomebaseUsers', 'toaster',
  function ($rootScope, $scope, $state, HomebaseUsers, toaster) {
    console.log("UsersCtrl] init");


    $scope.selectedUserCount = 0;
    $scope.userGridApi = null;
    $scope.usersGridOptions = {
      data: [],
      enableFullRowSelection: true,
      enableRowSelection: true,
      enableSelectAll: true,
      modifierKeysToMultiSelect: true,
      selectionRowHeaderWidth: 35,
      rowHeight: 35,
      showGridFooter:false,
      columnDefs: [{
        displayName: "ID",
        field: "_id",
        maxWidth: 220
      },{
        displayName: "Name",
        field: "name"
      },{
        displayName: "Username",
        field: "username"
      },{
        displayName: "Email",
        field: "email"
      },{
        displayName: "Created",
        field: "created"
      },{
        displayName: "Updated",
        field: "updated"
      }]
    };

    function loadTableData(){
      // Fetch table data
      HomebaseUsers.getAll().then(function(res){
        $scope.usersGridOptions.data = res;
      }).catch(function(err){
        toaster.pop('error', "Failed to fetch user list", err);
      })
    }

    $scope.updateSelectedCount = function(){ $scope.selectedUserCount = $scope.userGridApi.selection.getSelectedCount(); };

    $scope.usersGridOptions.onRegisterApi = function(gridApi){
      $scope.userGridApi = gridApi;
      gridApi.selection.on.rowSelectionChanged($scope, $scope.updateSelectedCount);
      gridApi.selection.on.rowSelectionChangedBatch($scope, $scope.updateSelectedCount);
    };

    $scope.editUser = function(){
      var currentSelection = $scope.userGridApi.selection.getSelectedRows()[0];
      $state.transitionTo("app.users.edit", {
        userId: currentSelection._id
      });
    };

    $scope.removeUser = function(){
      var currentSelection = $scope.userGridApi.selection.getSelectedRows()[0];
      HomebaseUsers.remove(currentSelection._id)
        .then(function(){ toaster.pop('success', "User Removed", "The user was removed successfully."); })
        .then(loadTableData)
        .catch(function(err){
          toaster.pop('error', "Failed to update user", err);
        })
    };

    // Generate a new height for the user table, capture a jQuery reference once during init
    var $tableContainer = $("#usersTableContainer");
    $scope.getTableHeight = function() { return { height: ($tableContainer.height() - 2) + "px" }; };

    loadTableData();
  }
]);

/**
 * Create User
 */
module.controller("CreateUserCtrl", ['$rootScope', '$scope', "$state", 'toaster', 'HomebaseUsers',
  function ($rootScope, $scope, $state, toaster, HomebaseUsers) {

    $scope.newUser = {
      name: "",
      username: "",
      email: "",
      password: ""
    };

    $scope.create = function(newUser){
      HomebaseUsers.create(newUser)
        .then(function(newUser){
          toaster.pop('success', "User Created", "New user was created successfully.");
          $state.transitionTo("app.users");
        })
        .catch(function(err){
          toaster.pop('error', "Failed to create new user", err);
        })
    }

  }
]);

/**
 * Edit User
 */
module.controller("EditUserCtrl", ['$rootScope', "$stateParams", '$scope', "$state", 'toaster', 'HomebaseUsers',
  function ($rootScope, $stateParams, $scope, $state, toaster, HomebaseUsers) {

    var editUserId = $stateParams.userId;

    HomebaseUsers.get($stateParams.userId)
      .then(function(userDef){
        $scope.user = userDef;
      })
      .catch(function(err){
        console.error("Cannot get userId: " + $stateParams.userId, err);
        editUserId = null;
      });


    $scope.update = function(updatedUser){
      HomebaseUsers.update(updatedUser)
        .then(function(updatedUser){
          toaster.pop('success', "User Updated", "User was updated successfully.");
          $state.transitionTo("app.users");
        })
        .catch(function(err){
          toaster.pop('error', "Failed to update user", err);
        })
    }

  }
]);
