// Create a Promise that waits for a condition.
function pollAndWait(timeout, check) {
  var INTERVAL = 100;
  return new Promise(function(resolve, reject) {
    var ival;
    ival = setInterval(function() {
      if (check()) {
        clearInterval(ival);
        resolve();
      }
      timeout -= INTERVAL;
      if (timeout < 0) {
        clearInterval(ival);
        reject('timeout surpassed');
      }
    }, INTERVAL);
  });
}
