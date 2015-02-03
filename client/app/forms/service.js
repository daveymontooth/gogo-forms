(function(){

    'use strict';

    angular.module('goGoForms')
        .service('formsService', Forms);

    Forms.$inject = ['$http','formsEndpoint'];

    function Forms($http,formsEndpoint){
        var svc = this;

        /* I like RESTful services, so I prefer to keep them light and stick
           to RESTful methods. (GET, PUT, POST, DELETE)
         */
        svc.get = function(id){
            formsEndpoint += (id) ? id : "";

            return $http.get(formsEndpoint)
                .success(function(data, status, headers){
                    return data;
                })
                .error(function(data, status, headers){
                    /* Logging the error for now */
                    console.log(status);
                    return {};
                });

        };
    }

})();
