(function(){

    'use strict';

    angular.module('goGoForms')
        .directive('formEdit', ['formsService', Edit]);

    function Edit(formsService){
        var directive = {
            restrict:'AE',
            link: Link,
            templateUrl:'app/forms/template.edit.html'
        };

        return directive;

        function Link(scope, element, attrs){
            var id = attrs && attrs.formid || null;
            formsService.get(id).then(function(response){
                scope.form = response.data;
                console.log(scope.form);
            });
        }
    }
})();
