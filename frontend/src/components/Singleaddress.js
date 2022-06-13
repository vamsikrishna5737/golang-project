import React ,{useState}from 'react'
import { useStateValue } from '../context/StateProvider'
import { actionType } from '../context/reducer'
import { useNavigate } from 'react-router-dom'

const Singleaddress = ({obj,fetchData}) => {
    const [state,dispatch] = useStateValue()
    const navigate=useNavigate()
    const edituser = async () =>{
      dispatch({type : actionType.ADD_Id, payload : {id:obj.id}})
      navigate("/form")
      }
    const deleteproduct = async (id) =>{
        
        const jsonData = await fetch(process.env.REACT_APP_API + "/delete/"+id , {
            method: "POST",
            headers: {
              "Content-Type": "application/json",
            },
          });
        const res = await jsonData.json()
        console.log(res)
        if (res.message === "success"){
            fetchData()
            navigate("/product")
        }  
    } 

  return (
    <tr>
                <th>{obj.productname}</th>
                <th>{obj.cost}</th>
                <th>{obj.mail}</th>
                <th><button onClick={()=>edituser()}>edit</button></th>
                <th><button onClick={()=>deleteproduct(obj.id)}>delete</button></th>
              </tr>
  )
}

export default Singleaddress