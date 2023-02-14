export interface Person {
  id: string;
  name: string;
  score: number;
}

export const fetchPeople = async (): Promise<Person[]> => {
  const res = await fetch("/api/");
  return res.json();
};

export const addPerson = async (person: Omit<Person, "id">): Promise<Response> => {
  return fetch("/api", {
    method: "POST",
    body: JSON.stringify(person),
    headers: {
      "Content-Type": "application/json",
    }
  });
};

export const updatePerson = async (person: Person): Promise<Response> => {
  return fetch("/api", {
    method: "PUT",
    body: JSON.stringify(person),
    headers: {
      "Content-Type": "application/json",
    }
  });
};

export const deletePerson = async (id: string): Promise<Response> => {
  return fetch(`/api/${id}`, {
    method: "DELETE"
  });
};
