import React,{useState} from 'react'
import { useStateValue } from '../context/StateProvider'
import { useNavigate } from 'react-router-dom'

const Form = () => {
    const [state,dispatch] = useStateValue()
    const [val,setVal] = useState("")
    const [val1,setVal1] = useState(0)
    const navigate=useNavigate()


    const handlesubmit = async (e) =>{
        e.preventDefault()
        const data={
            productname:val,
            cost:val1,
        }
        const jsonData = await fetch(process.env.REACT_APP_API + "/edit/"+state.id , {
            method: "POST",
            headers: {
              "Content-Type": "application/json",
            },
            body: JSON.stringify(data),
          });
        const res = await jsonData.json()
        console.log(res)
        if (res.message === "successfully updated"){
            navigate("/product")
        }  
        console.log(state.id)
    }
  return (
    <form onSubmit={handlesubmit}>
        <input type="text" value={val} onChange={(e)=>{setVal(e.target.value)}}/>
        <input type="number" value={val1} onChange={(e)=>{setVal1(e.target.value)}}/>
        <button type='submit'>submit</button>

    </form>
  )
}

export default Form