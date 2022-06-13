import React, { useEffect } from 'react'
import { useDispatch, useSelector } from 'react-redux'
import { deleteContact, fetchAll } from '../store/action/contactAction'
import { useNavigate } from 'react-router-dom';

const Home = () => {
  const dispatch = useDispatch()
  const navigate = useNavigate()
  const { contacts } = useSelector((state) => state.contactReducer)
  // console.log(contacts)
  useEffect(() => {
    dispatch(fetchAll())
  }, [])
  const detailHandler = (id) => {
    navigate(`/contacts/${id}`)
  }
  const deleteHandler = (id) => {
    dispatch(deleteContact(id))
      .then(data => dispatch(fetchAll()))
  }
  const addHandler = () => {
    navigate('/addcontact')
  }


  return (
    <div className='container'>
      <div className="card">
        <div className="card-header">
          <button className="btn btn-primary mx-1" onClick={addHandler}>Add Contact</button>
        </div>
        {contacts.map(el => {
          return <div className="card" key={el.ID}>
            <div className="card-body">
              <h5 className="card-title">{el.Name}</h5>
              <h6 className="card-subtitle mb-2 text-muted">{el.PhoneNumber}</h6>
              <div>
                <button className='btn btn-secondary mx-1' onClick={() => detailHandler(el.ID)}> See Detail</button>
                <button className='btn btn-danger' onClick={() => deleteHandler(el.ID)}>Delete</button>
              </div>
            </div>
          </div>
        })}
      </div>
    </div>
  )
}

export default Home