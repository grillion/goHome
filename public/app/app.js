var app = angular.module('app', [
  'ngAnimate',
  'ngRoute',
  "ui.bootstrap",
  "ui.layout",
  "ui.grid",
  "ui.grid.autoResize",
  "ui.grid.selection",
  "ui.router",
  "ncy-angular-breadcrumb",
  "toaster",
  "homebase-api",
  "homebase-api.auth",
  "homebase-api.devices",
  "homebase-api.settings",
  "homebase-api.users",
  "app.devices",
  "app.users"
]);

app.config(['$stateProvider', '$urlRouterProvider', function ($stateProvider) {

  console.log("App.config");

  $stateProvider
  // Default route, nothing specified check for login and re-route
    .state("app", {
      abstract: true,
      views: {
        "root": {
          controller: "LayoutCtrl",
          templateUrl: 'app/views/layout.html'
        }
      }
    })

    // Main page for non-authenticated users
    .state("login", {
      url: "/login",
      views: {
        root: {
          controller: "LoginCtrl",
          templateUrl: "app/views/login.html"
        }
      }
    })

    // Logout route
    .state("app.logout", {
      url: "/logout",
      controller: "LogoutCtrl",
      template: ""
    })

    // Main page for logged in users
    .state("app.home", {
      url: "/home",
      views: {
        "content@app": {
          controller: "HomeCtrl",
          templateUrl: "app/views/home.html"
        }
      },
      ncyBreadcrumb: {
        label: 'Home'
      }
    })

    // Device Configuration
    .state("app.devices", {
      url: "/devices",
      views: {
        "content@app": {
          controller: "DeviceListCtrl",
          templateUrl: "app/views/device/deviceList.html"
        }
      },
      ncyBreadcrumb: {
        label: 'Manage Devices'
      }
    })

    .state("app.devices.create", {
      url: "/create",
      views: {
        "content@app": {
          controller: "CreateDeviceCtrl",
          templateUrl: "app/views/device/createDevice.html"
        }
      },
      ncyBreadcrumb: {
        label: 'Create'
      }
    })

    .state("app.devices.edit", {
      url: "/edit/:deviceId",
      params: {
        deviceId: null
      },
      views: {
        "content@app": {
          controller: "EditDeviceCtrl",
          templateUrl: "app/views/device/editDevice.html"
        }
      },
      ncyBreadcrumb: {
        label: 'Edit'
      }
    })

    // System Settings
    .state("app.system", {
      url: "/system",
      views: {
        "content@app": {
          controller: "SystemCtrl",
          templateUrl: "app/views/system.html"
        }
      },
      ncyBreadcrumb: {
        label: 'System Settings'
      }
    })

    // Manage Users
    .state("app.users", {
      url: "/users",
      views: {
        "content@app": {
          controller: "UserListCtrl",
          templateUrl: "app/views/user/userList.html"
        }
      },
      ncyBreadcrumb: {
        label: 'Manage Users'
      }
    })

    .state("app.users.create", {
      url: "/create",
      views: {
        "content@app": {
          controller: "CreateUserCtrl",
          templateUrl: "app/views/user/createUser.html"
        }
      },
      ncyBreadcrumb: {
        label: 'Create'
      }
    })

    .state("app.users.edit", {
      url: "/edit/:userId",
      params: {
        userId: null
      },
      views: {
        "content@app": {
          controller: "EditUserCtrl",
          templateUrl: "app/views/user/editUser.html"
        }
      },
      ncyBreadcrumb: {
        label: 'Edit'
      }
    })

}]);

/**
 * App Init
 */
app.run(['$rootScope', '$state', function($rootScope, $state) {

  console.log("App.run");

  $rootScope.loggedIn = false;
  $rootScope.pageTitle = "";

  //// For each state change, if the user isn't logged in go to login
  $rootScope.$on("$stateChangeStart", function(event, toState, toParams, fromState, fromParams) {
    console.log("State changed", toState);
    //if( !$rootScope.loggedIn && toState.name != "app.login"){
    //  $state.transitionTo("app.login");
    //}
    //
    //if( $rootScope.loggedIn && next.$$route.originalPath == "/"){
    //  $state.transitionTo("app.home"
    //}
  });

}]);

app.directive('hbNav', function() {
  return {
    restrict: "E",
    replace: true,
    templateUrl: "app/views/nav.html",
    link: function(scope, element) {
      console.log("hbNav link");
    }
  };
});

/**
 * Default route / Layout Controller
 */
app.controller("LayoutCtrl", ['$rootScope', '$scope', function ($rootScope, $scope) {
  console.log("LayoutCtrl] Layout initialized");

  $scope.logout = function(){
    $rootScope.loggedIn = false;
  }
}]);

/**
 * Login Route
 **/
app.controller("LoginCtrl", [ '$rootScope', '$scope', '$state', 'HomebaseAuth',
  function ($rootScope, $scope, $state, HomebaseAuth) {
  console.log("LoginCtrl] init");
  $rootScope.pageTitle = "Login";
  $scope.loginError = false;
  $scope.loginUsername = "";
  $scope.loginPass = "";
  $scope.submitLogin = function($event){
    console.log("Login submission handler", $event);

    HomebaseAuth.login($scope.loginUsername, $scope.loginPass)
      .then(function(result){
        $rootScope.loggedIn = true;
        $rootScope.username = result.username;
        $state.transitionTo("app.home");
      })
      .catch(function(err){
        console.error(err);
        $scope.loginError = err;
      });

  };
}]);

/**
 * Logout Route
 */
app.controller("LogoutCtrl", ['$rootScope', '$scope', '$location', function ($rootScope, $scope, $location) {
  console.log("LogoutCtrl] init");
  $rootScope.loggedIn = false;
  $rootScope.username = "";
  $location.path("/");

  $scope.$on('ui.layout.toggle', function(e, container){
    console.log("Container toggled", container);
    //if ( container.size > 0 ){ console.log('container is open!'); }
  });
}]);

/**
 * Default route
 */
app.controller("HomeCtrl", ['$scope', function ($scope) {
  console.log("HomeCtrl] init");
}]);

/**
 * Settings route
 */
app.controller("SystemCtrl", ['$rootScope', '$scope', 'HomebaseUsers', function ($rootScope, $scope, HomebaseUsers) {
  console.log("SystemCtrl] init");
}]);

console.log("Controllers defined");