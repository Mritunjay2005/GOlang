const token = () => localStorage.getItem('token');
const headers = () => ({
  'Content-Type':'application/json',
  'Authorization': `Bearer ${token()}`
});

function logout() {
  localStorage.removeItem('token');
  window.location.href = 'index.html';
}

async function loadExpenses() {
  const cat  = document.getElementById('filter-cat').value;
  const from = document.getElementById('filter-from').value;
  const to   = document.getElementById('filter-to').value;
  let url = '/api/expenses?';
  if (cat)  url += `category=${cat}&`;
  if (from) url += `from=${from}&`;
  if (to)   url += `to=${to}&`;

  const res = await fetch(url, { headers: headers() });
  if (res.status === 401) { logout(); return; }
  const expenses = await res.json() || [];
  renderTable(expenses);
  window._expenses = expenses;
  updateCharts(expenses);
  loadSummary();
}

function renderTable(expenses) {
  const tbody = document.getElementById('expense-body');
  tbody.innerHTML = '';
  expenses.forEach(e => {
    const tr = document.createElement('tr');
    tr.innerHTML = `
      <td>${e.expense_date}</td>
      <td>${e.title}</td>
      <td>${e.category}</td>
      <td>₹${parseFloat(e.amount).toFixed(2)}</td>
      <td>${e.note || '—'}</td>
      <td>
        <button class="btn-edit"   onclick="editExpense(${e.id})">Edit</button>
        <button class="btn-delete" onclick="deleteExpense(${e.id})">Del</button>
      </td>`;
    tbody.appendChild(tr);
  });
}

async function saveExpense() {
  const id       = document.getElementById('f-id').value;
  const payload  = {
    title:        document.getElementById('f-title').value,
    amount:       parseFloat(document.getElementById('f-amount').value),
    category:     document.getElementById('f-category').value,
    note:         document.getElementById('f-note').value,
    expense_date: document.getElementById('f-date').value,
  };
  const url    = id ? `/api/expenses/${id}` : '/api/expenses';
  const method = id ? 'PUT' : 'POST';
  await fetch(url, { method, headers: headers(), body: JSON.stringify(payload) });
  clearForm();
  loadExpenses();
}

function editExpense(id) {
  const e = (window._expenses || []).find(x => x.id === id);
  if (!e) return;
  document.getElementById('f-id').value       = e.id;
  document.getElementById('f-title').value    = e.title;
  document.getElementById('f-amount').value   = e.amount;
  document.getElementById('f-category').value = e.category;
  document.getElementById('f-date').value     = e.expense_date;
  document.getElementById('f-note').value     = e.note || '';
}

async function deleteExpense(id) {
  if (!confirm('Delete this expense?')) return;
  await fetch(`/api/expenses/${id}`, { method:'DELETE', headers: headers() });
  loadExpenses();
}

function clearForm() {
  ['f-id','f-title','f-amount','f-note'].forEach(i => document.getElementById(i).value='');
  document.getElementById('f-category').value = '';
  document.getElementById('f-date').value     = '';
}

async function loadSummary() {
  const res = await fetch('/api/expenses/summary', { headers: headers() });
  const s   = await res.json();
  document.getElementById('total-amount').textContent = `₹${parseFloat(s.total).toFixed(2)}`;
  const cats   = Object.entries(s.by_category || {});
  const topCat = cats.sort((a,b)=>b[1]-a[1])[0];
  document.getElementById('top-cat').textContent = topCat ? topCat[0] : '—';
}

// Init
if (!token()) window.location.href = 'index.html';
loadExpenses();