import React, { useEffect, useState } from 'react'
import { useDispatch, useSelector } from 'react-redux'
import { fetchContact } from '../store/action/contactAction'
import { useNavigate } from 'react-router-dom';
import { addContact } from '../store/action/contactAction';
import swal from "sweetalert"

const Form = () => {
  const dispatch = useDispatch()
  const navigate = useNavigate()
  const [name, setName] = useState("")
  const [phoneNumber, setPhoneNumber] = useState("")
  const [email, setEmail] = useState("")

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

  const homeHandler = () => {
    navigate('/')
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
      dispatch(addContact(payload))
        .then(res => {
          navigate('/')
        })
    }
  }

  return (
    <div className='container' >
      <form onSubmit={submitHandler}>
        <h4>Add a Contact</h4>
        <div className="form-group">
          <label htmlFor="name">Name</label>
          <input type="text" className="form-control" id="name" aria-describedby="emailHelp" onChange={nameHandler} />
        </div>
        <div className="form-group">
          <label htmlFor="phoneNumber">Phone Number</label>
          <input type="text" className="form-control" id="PhoneNumber"
            onChange={phoneNumberHandler} />
        </div>
        <div className="form-group">
          <label htmlFor="email">Email</label>
          <input type="email" className="form-control" id="email"
            onChange={emailHandler} />
        </div>
        <button type='button' className="btn btn-secondary mx-2" onClick={homeHandler}>Back to Home</button>
        <button type="submit" className="btn btn-primary">Submit</button>
      </form>
    </div >
  )
}

export default Form