(function(){

    'use strict';

    angular.module('goGoForms', ['ui.router'])
        .config(function($provide, $stateProvider, $urlRouterProvider){
            $provide.value('formsEndpoint', '<Put your forms endpoint here>');

            $urlRouterProvider.otherwise("/");

            $stateProvider
                .state('home', {
                    url:'/',
                    templateUrl:'assets/states/forms/list.html'
                })
                .state('create',{
                    url:'/create',
                    templateUrl:'assets/states/forms/create.html'
                })
                .state('edit',{
                    url:'/edit/:id',
                    templateUrl:'assets/states/forms/edit.html',
                    controller: function($scope, $stateParams) {
                        $scope.id = $stateParams.id;
                    }
                })
        });

})();
