import {create} from "zustand"

export type EditedTask = {
    id : number
    title :string
}

type State = {
    editedTask : EditedTask,
    updateEditedTask : (payload : EditedTask) => void
    resetEditedTask : () => void
}

const useStore = create((set) => ({
    editedTask : {id : 0 , title : ""},
    updatedEditedTask : (payload : EditedTask) =>
        set({
            editedTask : payload,
        }),
    resetEditedTask : () => set({editedTask : {id :0 , title : ""}}),
}))

export default useStore