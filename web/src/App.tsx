import React, { Fragment, useEffect } from 'react';
import { useMutation, useQuery } from 'react-query';
import { addPerson, deletePerson, fetchPeople, Person } from './api/Person';
import './App.css';

function App() {
  const { data, refetch } = useQuery("people", fetchPeople);

  const createPerson = useMutation((person: Omit<Person, "id">) => {
    return addPerson(person);
  });

  const removePerson = useMutation((id: string) => {
    return deletePerson(id)
  });

  useEffect(() => {
    if (removePerson.isSuccess) {
      refetch();
    }
  }, [removePerson.isSuccess]);

  useEffect(() => {
    if (createPerson.isSuccess) {
      refetch();
    }
  }, [createPerson.isSuccess]);

  return (
    <div>
      <div className="grid">
        <p>Name</p>
        <p>Score</p>
        <p>Options</p>
        {
          data ? data.map(({ id, name, score }) =>
            <Fragment key={name}>
              <p>{name}</p>
              <p>{score}</p>
              <button className="btn" onClick={() => removePerson.mutate(id)}>Delete</button>
            </Fragment>
          ) : <p className="loading">Loading data</p>
        }
      </div>
      <hr />
      <form className="add-person" onSubmit={(e: any) => {
        e.preventDefault();

        const name = e.target[0].value;
        const score = parseFloat(e.target[1].value);

        createPerson.mutate({ name, score });
      }}>
        Name: <input />
        Score: <input type="number" />
        <button type="submit">Add Person</button>
      </form>
    </div>
  );
}

export default App;
