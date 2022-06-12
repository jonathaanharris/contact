import { CONTACT_FETCH_SUCCESS } from "./type";

export const fetchAll = () => {
  return (dispatch) => {
    return fetch(`http://localhost:8080/contact`)
      .then(res => {
        if (!res.ok) {
          throw new Error('something went wrong')
        }
        return res.json()
      })
      .then(data => {
        console.log(data)

        return dispatch(fetchContactSuccess(data))
      })
  }
}

export const updateContact = (data, id) => {
  return (dispatch) => {
    console.log(JSON.stringify(data), id, 388888)
    return fetch(`http://localhost:8080/contact/${id}`, {
      method: "PATCH",
      headers: {
        'Content-Type': 'application/json',
      },
      body: JSON.stringify(data)
    })
      .then(res => {
        if (!res.ok) {
          throw new Error('something went wrong')
        }
        return res.json
      })
      .then(data => {
        return data
      })
  }
}
export const addContact = (data) => {
  return (dispatch) => {
    console.log(JSON.stringify(data))
    return fetch('http://localhost:8080/contact', {
      method: "POST",
      headers: {
        'Content-Type': 'application/json',

      },
      body: JSON.stringify(data)
    })
      .then(res => {
        if (!res.ok) {
          throw new Error('something went wrong')
        }
        return res.json
      })
      .then(data => {

        return data
      })
  }
}

export const deleteContact = (data) => {
  return (dispatch) => {
    return fetch(`http://localhost:8080/contact/${data}`, {
      method: "DELETE"
    })
      .then(res => {
        if (!res.ok) {
          throw new Error('something went wrong')
        }
        return res.json
      })
      .then(data => data)
  }
}

export const fetchContactSuccess = (data) => {
  return {
    type: CONTACT_FETCH_SUCCESS,
    data
  }
}