export const actionType = {
  ADD_USER: "add user",
  ADD_Token: "add token",
  ADD_Products: "add products",
  ADD_Id: "add id",
};

const reducer = (state, action) => {
  switch (action.type) {
    case actionType.ADD_USER:
      return { ...state, user: action.payload.user };

    case actionType.ADD_Token:
      return {
        ...state,
        token: action.payload.token,
      };
    case actionType.ADD_Products:
      return { ...state, products: action.payload.products}  

    case actionType.ADD_Id:
      return { ...state, id: action.payload.id}

    default:
      return state;
  }
};

export default reducer;
