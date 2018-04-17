let expenses = document.getElementById('expenses');
let expenseName = document.getElementById('expense-name');
let expenseIncome = document.getElementById('expense-income');
let expenseButton = document.getElementById('expense-button');
let userBalance = document.getElementById('balance');
function newExpense() {
  var AJAX = new XMLHttpRequest();
  AJAX.open('POST', '/new', true);
  AJAX.setRequestHeader('Content-Type', 'application/x-www-form-urlencoded');
  AJAX.send('expense-name=' + expenseName.value + '&expense-income=' + expenseIncome.value);
  AJAX.onreadystatechange = () => {
    if (AJAX.readyState == 4 && AJAX.status == 200) {
      var res = JSON.parse(AJAX.responseText);
      expenses.innerHTML = `<div id="${res.uuid}" class="expense"><span class="expense-name">${expenseName.value}</span><span class="expense-time">${res.time}</span><span class="expense-income">${expenseIncome.value} 円</span><input class="expense-erase" type="button" onclick="eraseExpense("${res.uuid}");"></div>` + expenses.innerHTML;
    }
  }
}
function eraseExpense(uuid) {
  if (confirm('本当に削除してもいいですか？')) {
    var AJAX = new XMLHttpRequest();
    AJAX.open('POST', '/erase', true);
    AJAX.setRequestHeader('Content-Type', 'application/x-www-form-urlencoded');
    AJAX.send('expense-uuid=' + uuid);
    AJAX.onreadystatechange = () => {
      if (AJAX.readyState == 4 && AJAX.status == 200) {
        document.getElementById(uuid).remove();
        userBalance.innerText = AJAX.responseText + ' 円';
      }
    }
  }
}
