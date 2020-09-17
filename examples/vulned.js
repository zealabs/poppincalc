var exec = require('child_process').exec;

module.exports = {
    execShit: function (cmd) {
      console.log(cmd);
      exec(cmd);
    }
  };
  
var main = function () {
    console.log('Vulnerable exported function! but main func not called.');
}