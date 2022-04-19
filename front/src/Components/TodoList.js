import PropTypes from 'prop-types';
import { TodoItem } from './TodoItem';
import '../css/TodoList.scss';
// import { useState } from 'react';

const HOST_ADDRESS = 'http://localhost:8080';

export const TodoList = ({ todos }) => {
  const handleRemove = (id) => {
    const url = HOST_ADDRESS + '/delete/' + id;
    fetch(url, {
      method: 'DELETE',
      headers: {
        'Content-Type': 'application/json',
      },
    })
      .then((response) => response.json())
      .catch((error) => console.log(error));
  };

  return (
    <ul className="todo-list">
      {todos.map((todoItem) => (
        <li key={todoItem.id}>
          <TodoItem
            title={todoItem.title}
            // description={todoItem.description}
            completed={todoItem.completed}
            id={todoItem.id}
            handleRemove={() => handleRemove(todoItem.id)}
          />
        </li>
      ))}
    </ul>
  );
};

TodoItem.propTypes = {
  todos: PropTypes.arrayOf(PropTypes.object),
};
