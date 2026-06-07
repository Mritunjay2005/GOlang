const API = '';   // same origin

function showTab(tab) {
  document.getElementById('login-section').hidden    = tab !== 'login';
  document.getElementById('register-section').hidden = tab !== 'register';
  document.querySelectorAll('.tab').forEach((b,i)=>
    b.classList.toggle('active', (i===0&&tab==='login')||(i===1&&tab==='register')));
}

async function login() {
  const email    = document.getElementById('login-email').value;
  const password = document.getElementById('login-password').value;
  const res = await fetch(`${API}/api/login`, {
    method:'POST', headers:{'Content-Type':'application/json'},
    body: JSON.stringify({email, password})
  });
  if (!res.ok) {
    document.getElementById('login-error').textContent = 'Invalid credentials';
    return;
  }
  const data = await res.json();
  localStorage.setItem('token', data.token);
  window.location.href = 'dashboard.html';
}

async function register() {
  const name     = document.getElementById('reg-name').value;
  const email    = document.getElementById('reg-email').value;
  const password = document.getElementById('reg-password').value;
  const res = await fetch(`${API}/api/register`, {
    method:'POST', headers:{'Content-Type':'application/json'},
    body: JSON.stringify({name, email, password})
  });
  if (!res.ok) {
    document.getElementById('reg-error').textContent = 'Registration failed (email taken?)';
    return;
  }
  showTab('login');
}

// Redirect if already logged in
if (localStorage.getItem('token')) window.location.href = 'dashboard.html';