import React, { useEffect} from "react";
import { useStateValue } from "../context/StateProvider";
import { actionType } from "../context/reducer";
import Header from "../components/Header";
import Aside from "../components/Aside";
import AllProducts from "../components/AllProducts";

const TotalProducts = () => {
  const [state, dispatch] = useStateValue();

  const user = async ()=>{

    // console.log(state.token)
    const prores = await fetch(process.env.REACT_APP_API + "/auth/" + state.token, {
      method : "POST",
      headers: { "Content-Type": "application/json" },
    })

    const user = await prores.json()
    console.log(user)
    if (user.message ==="success"){
      const user1 = user.token
      return dispatch({ type: actionType.ADD_USER, payload: { user: user1}})
    }
    return 
  }
  useEffect(()=>{
    user()
  },[]);

  return (
    <>
      <div className="mailContainer">
        <Aside/>
        <div className="innerContainer">
          <Header/>
          <AllProducts />
        </div>
      </div>
      </>
  );
};

export default TotalProducts;