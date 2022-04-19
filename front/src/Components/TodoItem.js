import PropTypes from 'prop-types';
import React from 'react';
import '../css/TodoItem.scss';

const HOST_ADDRESS = 'http://localhost:8080';

export const TodoItem = ({
  id,
  title,
  description,
  completed,
  handleRemove,
}) => {
  const handleClick = (e) => {
    const url = HOST_ADDRESS + '/put/' + id;
    fetch(url, {
      method: 'PUT',
      headers: {
        'Content-Type': 'application/json',
      },
    })
      .then((response) => response.json())
      .catch((error) => console.log(error));
  };

  return (
    <article className="todo-item">
      <section className="todo-item-text">
        <h3> {title}</h3>
        <p> {description} </p>
      </section>
      <button onClick={handleRemove}> Delete </button>
      <input onClick={handleClick} type="checkbox" defaultChecked={completed} />
    </article>
  );
};

TodoItem.propTypes = {
  title: PropTypes.string.isRequired,
  description: PropTypes.string,
  completed: PropTypes.bool.isRequired,
  handleRemove: PropTypes.func.isRequired,
  id: PropTypes.string.isRequired,
};
