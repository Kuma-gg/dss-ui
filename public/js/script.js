(function () {
    var modals = document.querySelectorAll('.modal');
    var datepickers = document.querySelectorAll('.datepicker');
    var navs = document.querySelectorAll('.sidenav');
    var tabs = document.getElementById('tabs');

    M.Modal.init(modals);
    M.Sidenav.init(navs, {});
    M.Tabs.init(tabs)
    M.Datepicker.init(datepickers, { format: 'dd/mm/yyyy', maxDate: new Date(), container: document.body });
})();