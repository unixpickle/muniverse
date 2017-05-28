// Stubs to make Famobi games work correctly.

var _fgq = [];

function fg_api(data) {
  this._data = data;
  this.localStorage = window.localStorage;
  localStorage.clear();
}

fg_api.EMPTY_IMAGE = 'data:image/gif;base64,R0lGODlhAQABAAAAACH5BAEKAAEALAAAAAABAAEAAAICTAEAOw==';

fg_api.prototype.__ = function(key) {
  if (key === 'preload_image') {
    return fg_api.EMPTY_IMAGE;
  }
  var res = this._data.game_i18n.en[key] ||
    this._data.game_i18n.default[key] ||
    this._data.branding[key];
  return res.replace('{lang}', 'en');
}

fg_api.prototype.getMoreGamesButtonImage = function() {
  // Empty GIF, https://stackoverflow.com/questions/9126105/blank-image-encoded-as-data-uri.
  return fg_api.EMPTY_IMAGE;
};

fg_api.prototype.moreGamesLink = function() {
  return '';
};

fg_api.prototype.submitHighscore = function() {
};

fg_api.prototype.showAd = function() {
};
