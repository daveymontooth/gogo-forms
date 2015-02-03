module.exports = function(grunt) {
    grunt.initConfig({
        pkg: grunt.file.readJSON('node_modules/grunt/package.json'),
        sass: {
            dist: {
                options: {
                    style:'compressed'
                },
                files: {
                    'assets/css/style.css' : 'assets/sass/style.scss'
                }
            }
        },
        watch: {
            css: {
                files: '**/*.scss',
                tasks: ['sass']
            }
        }
    });
    grunt.loadNpmTasks('grunt-contrib-sass');
    grunt.loadNpmTasks('grunt-contrib-watch');
    grunt.registerTask('default',['watch']);
}
