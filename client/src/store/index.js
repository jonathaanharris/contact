import { createStore, combineReducers, applyMiddleware } from "redux"
import thunk from "redux-thunk"
import contactReducer from "./reducer/contactReducer"

const rootReducer = combineReducers({
  contactReducer
})

let store = createStore(rootReducer, applyMiddleware(thunk))

export default store