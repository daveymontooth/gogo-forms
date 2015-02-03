(function(){

    'use strict';

    angular.module('goGoForms')
        .directive('formsList', ['formsService', List]);

    function List(formsService){

         var directive = {
             restrict:'AE',
             link: Link,
             templateUrl:'app/forms/template.list.html'
         };

        return directive;

        function Link(scope, element, attrs){

            formsService.get().then(function(response){
                scope.forms = response.data;
            });
        }
    }
})();
