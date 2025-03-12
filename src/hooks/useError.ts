import {useNavigate} from "react-router-dom";
import {CrsfToken} from "../types"
import useStore from "../store"
import axios from "axios";

export const useError = () => {
    const navigate = useNavigate()
    const resetEditedTask = useStore((state) => state.resetEditedTask))
    const getCsrfToken = async () => {
        const {data} = await axios.get<CrsfToken>(
            `${process.env.REACT_APP_API_URL}/csrf`
        )
        axios.defaults.headers.common["X-CSRF-TOKEN"] = data.csrf_token
    }
    const switchErrorHandling = (msg : string) => {
        switch (msg){
            case "invalid csrf token":
                getCsrfToken()
                alert("CSRFトークンがありません、もう一度アクセスしなおしてください")
                break;

            case "invalid or expired jwt":
                alert("無効なトークンの有効期限が切れてください。ログインしてください")
                resetEditedTask()
                navigate("/")
                break
            case "missing or malformed jwt":
                alert(`アクセストークンが無効です、ログインしてください`)
                resetEditedTask()
                navigate("/")
                break
            case "duplicated key not allowed":
                alert("すでに電子メールが存在します。別のメールを使用してください")
                break
            case "crypto/bcrypt: hashedPassword is not the hash of the given password":
                alert(`パスワードが正しくないです`)
                break
            case "record not found":
                alert(`メールアドレスが正しくないです`)
                break
            default:
                alert(msg)
        }
    }
    return { switchErrorHandling}
}