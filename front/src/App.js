import React, { useState } from 'react';
import { Header } from './Components/Header';
import { TodoList } from './Components/TodoList';
import './css/index.css';

const HOST_ADDRESS = 'http://localhost:8080';

export const App = () => {
  const [todos, setTodos] = useState([]);
  const [todo, setTodo] = useState('');

  const handleSubmit = (e) => {
    // e.preventDefault();
    const url = HOST_ADDRESS + '/post/' + todo;
    fetch(url, {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json',
      },
      // body: JSON.stringify(data),
    })
      .then((response) => response.json())
      .catch((error) => console.log(error));
    setTodo('');
  };

  const update = async () => {
    setTodos([]);
    const result = await fetch(HOST_ADDRESS);
    const data = await result.json();
    setTodos(data);
  };

  return (
    <div className="App">
      <Header />
      <button onClick={update}> Refresh </button>
      <form onSubmit={handleSubmit}>
        <input
          type="text"
          value={todo}
          onChange={(e) => setTodo(e.target.value)}
        />
        <button>Submit</button>
      </form>
      <TodoList todos={todos} />
    </div>
  );
};

export default App;
