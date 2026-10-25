(() => {
  const messages = document.getElementById('messages');
  const form = document.getElementById('chat-form');
  const input = document.getElementById('message-input');

  let ws;

  function connect() {
    const protocol = location.protocol === 'https:' ? 'wss:' : 'ws:';
    ws = new WebSocket(`${protocol}//${location.host}/ws`);

    ws.onmessage = (event) => {
      const msg = document.createElement('div');
      msg.className = 'msg';
      msg.textContent = event.data;
      messages.appendChild(msg);
      messages.scrollTop = messages.scrollHeight;
    };

    ws.onclose = () => {
      setTimeout(connect, 1000);
    };
  }

  form.addEventListener('submit', (e) => {
    e.preventDefault();
    const text = input.value.trim();
    if (text && ws && ws.readyState === WebSocket.OPEN) {
      ws.send(text);
      input.value = '';
    }
  });

  connect();
})();
