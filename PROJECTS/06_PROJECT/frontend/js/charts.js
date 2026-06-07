let pieChart, barChart;

function updateCharts(expenses) {
  const catMap = {};
  expenses.forEach(e => {
    catMap[e.category] = (catMap[e.category] || 0) + parseFloat(e.amount);
  });
  const labels = Object.keys(catMap);
  const data   = Object.values(catMap);
  const colors = [
    '#6366f1','#10b981','#f59e0b','#ef4444',
    '#3b82f6','#8b5cf6','#14b8a6','#f97316'
  ];

  // Pie
  if (pieChart) pieChart.destroy();
  pieChart = new Chart(document.getElementById('pie-chart'), {
    type: 'pie',
    data: { labels, datasets:[{ data, backgroundColor: colors }] },
    options: { plugins:{ legend:{ position:'bottom' } } }
  });

  // Bar
  if (barChart) barChart.destroy();
  barChart = new Chart(document.getElementById('bar-chart'), {
    type: 'bar',
    data: {
      labels,
      datasets:[{ label:'Spent (₹)', data, backgroundColor: colors }]
    },
    options: { scales:{ y:{ beginAtZero:true } } }
  });
}