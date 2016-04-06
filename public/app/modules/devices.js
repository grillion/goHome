var module = angular.module("app.devices", []);

module.directive('deviceForm', function(){
  return {
    restrict: "E",
    replace: true,
    templateUrl: "app/views/device/deviceForm.html",
    scope: {
      device: "=",
      onSubmit: "="
    }
  };
});

/**
 * List Devices
 */
module.controller("DeviceListCtrl", ['$rootScope', '$scope', '$state', 'HomebaseDevices', 'toaster',
  function ($rootScope, $scope, $state, HomebaseDevices, toaster) {
    console.log("DeviceListCtrl] init");

    $scope.selectedDeviceCount = 0;
    $scope.deviceGridApi = null;
    $scope.deviceGridOptions = {
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
        displayName: "Type",
        field: "type",
        maxWidth: 100
      },{
        displayName: "Name",
        field: "name"
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
      HomebaseDevices.getAll().then(function(res){
        $scope.deviceGridOptions.data = res;
      }).catch(function(err){
        toaster.pop('error', "Failed to fetch device list", err);
      })
    }

    $scope.updateSelectedCount = function(){ $scope.selectedDeviceCount = $scope.deviceGridApi.selection.getSelectedCount(); };

    $scope.deviceGridOptions.onRegisterApi = function(gridApi){
      $scope.deviceGridApi = gridApi;
      gridApi.selection.on.rowSelectionChanged($scope, $scope.updateSelectedCount);
      gridApi.selection.on.rowSelectionChangedBatch($scope, $scope.updateSelectedCount);
    };

    $scope.editDevice = function(){
      var currentSelection = $scope.deviceGridApi.selection.getSelectedRows()[0];
      $state.transitionTo("app.devices.edit", {
        deviceId: currentSelection._id
      });
    };

    $scope.removeDevice = function(){
      var currentSelection = $scope.deviceGridApi.selection.getSelectedRows()[0];
      HomebaseDevices.remove(currentSelection._id)
        .then(function(){ toaster.pop('success', "Device Removed", "The device was removed successfully."); })
        .then(loadTableData)
        .catch(function(err){
          toaster.pop('error', "Failed to update device", err);
        })
    };

    // Generate a new height for the device table, capture a jQuery reference once during init
    var $tableContainer = $("#devicesTableContainer");
    $scope.getTableHeight = function() { return { height: ($tableContainer.height() - 2) + "px" }; };

    loadTableData();
  }
]);

/**
 * Create Device
 */
module.controller("CreateDeviceCtrl", ['$rootScope', '$scope', "$state", 'toaster', 'HomebaseDevices',
  function ($rootScope, $scope, $state, toaster, HomebaseDevices) {

    $scope.newDevice = {
      name: "",
      type: ""
    };

    $scope.create = function(createDevice){
      HomebaseDevices.create(createDevice)
        .then(function(newDevice){
          toaster.pop('success', "Device Created", "New device was created successfully.");
          $state.transitionTo("app.devices");
        })
        .catch(function(err){
          toaster.pop('error', "Failed to create new device", err);
        })
    }

  }
]);

/**
 * Edit Device
 */
module.controller("EditDeviceCtrl", ['$rootScope', "$stateParams", '$scope', "$state", 'toaster', 'HomebaseDevices',
  function ($rootScope, $stateParams, $scope, $state, toaster, HomebaseDevices) {

    var editDeviceId = $stateParams.deviceId;

    HomebaseDevices.get($stateParams.deviceId)
      .then(function(deviceDef){
        $scope.device = deviceDef;
      })
      .catch(function(err){
        console.error("Cannot get deviceId: " + $stateParams.deviceId, err);
        editDeviceId = null;
      });

    $scope.update = function(updateDevice){
      HomebaseDevices.update(updateDevice)
        .then(function(updatedDevice){
          toaster.pop('success', "Device Updated", "Device was updated successfully.");
          $state.transitionTo("app.devices");
        })
        .catch(function(err){
          toaster.pop('error', "Failed to update device", err);
        })
    }

  }
]);
