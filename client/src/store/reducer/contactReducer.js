import { CONTACT_FETCH_SUCCESS } from "../action/type"

const initialState = {
  contacts: []
}

function contactReducer(state = initialState, action) {
  if (action.type === CONTACT_FETCH_SUCCESS) {
    return {
      ...state, contacts: action.data
    }
  }
  return state
}

export default contactReducer