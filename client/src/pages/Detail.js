import React, { useEffect, useState } from 'react'
import { useSelector, useDispatch } from 'react-redux'
import { useParams, useNavigate } from 'react-router-dom'
import { fetchAll } from '../store/action/contactAction'
import swal from "sweetalert"
import { updateContact } from '../store/action/contactAction';

const Detail = () => {
  const { id } = useParams()
  const dispatch = useDispatch()
  const navigate = useNavigate()

  const { contacts } = useSelector((state) => state.contactReducer)

  const contact = contacts.find(el => {
    return el.ID == id
  })


  const [name, setName] = useState(contact?.Name ? contact.Name : "")
  const [phoneNumber, setPhoneNumber] = useState(contact?.PhoneNumber ? contact.PhoneNumber : "")
  const [email, setEmail] = useState(contact?.Email ? contact.Email : "")

  useEffect(() => {
    dispatch(fetchAll())
  }, [])

  const homeHandler = () => {
    navigate('/')
  }


  if (contacts.length === 0) {
    return <div>loading</div>
  }

  if (!contact) {
    return <div>
      <p class="text-uppercase">data not found</p>
      <button type='button' className="btn btn-secondary mx-2" onClick={homeHandler}>Back to Home</button>

    </div >

  }

  const nameHandler = (e) => {
    setName(e.target.value)
  }
  const phoneNumberHandler = (e) => {
    let temp = e.target.value
    if (e.target.value[0] == 0) {
      temp = "+62" + e.target.value.slice(1)
    }
    setPhoneNumber(String(temp))
  }
  const emailHandler = (e) => {
    setEmail(e.target.value)
  }


  const submitHandler = (e) => {
    e.preventDefault();
    let payload = { Name: name, PhoneNumber: phoneNumber, Email: email }
    if (!name) {
      swal("Name is required")
    } else if (!phoneNumber) {
      swal("phoneNumber is required")
    }
    else if (!email) {
      swal("email is required")
    }
    else if (phoneNumber[0] !== "0" && phoneNumber.slice(0, 3) !== "+62") {
      swal('wrong phone Number Format')
    } else if (!email.includes("@")) {
      swal('wrong email Format')
    }
    else {
      dispatch(updateContact(payload, id))
        .then(res => {
          navigate('/')
        })
    }
  }


  return (

    <div className='container' >
      <form onSubmit={submitHandler}>
        <h4>Update Contact</h4>
        <div className="form-group">
          <label htmlFor="name">Name</label>
          <input value={name} type="text" className="form-control" id="name" aria-describedby="emailHelp" onChange={nameHandler} />
        </div>
        <div className="form-group">
          <label htmlFor="phoneNumber">Phone Number</label>
          <input value={phoneNumber} type="text" className="form-control" id="PhoneNumber"
            onChange={phoneNumberHandler} />
        </div>
        <div className="form-group">
          <label htmlFor="email">Email</label>
          <input value={email} type="email" className="form-control" id="email"
            onChange={emailHandler} />
        </div>
        <button type='button' className="btn btn-secondary mx-2" onClick={homeHandler}>Back to Home</button>
        <button type="submit" className="btn btn-primary">Submit</button>
      </form>
    </div >
  )
}

export default Detail