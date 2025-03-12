import React, {FormEvent, useState} from 'react';
import {useMutateAuth} from "../hooks/useMutateAuth";

const Auth = () => {
    const [email, setEmail] = useState<string>("")
    const [pw, setPw] = useState<string>("")
    const [isLogin, setLogin] = useState<boolean>(true)
    const {loginMutation , registerMutation} = useMutateAuth()

    const submitAuthHandler = async (e : FormEvent<HTMLFormElement>) => {
        e.preventDefault()
        if (isLogin) {
            loginMutation.mutate({
                email :email,
                password : pw
            })
        }else {
            await registerMutation.mutateAsync({
                email : email,
                password : pw
            }).then(() => loginMutation.mutate({email : email , password : pw}))
        }
    }

    return (
        <div className="flex justify-content items-center flex-col min-h-screen text-gray-600 font-mono">
            <div className="flex items-center">
                {/*<CheckBadgeIcon className="h-8 w-8 mr-2 text-bule-500"/>*/}
                <span className="text-center text-3xl font-extrabold">
                    Todoアプリ　React / Go (Echo)
                </span>
            </div>
            <h2 className="my-6">{isLogin ? "ログイン" : "新規作成"}</h2>
            <form onSubmit={submitAuthHandler}>
                <div><input  className="mb-3 px-3 text-sm py-2 border border-gray-300"
                             onChange={(e) => setEmail(e.target.value)}
                             type="email" name={"email"} autoFocus placeholder={"メールアドレス"}/>
                </div>
                <div><input  className="mb-3 px-3 text-sm py-2 border border-gray-300"
                             onChange={(e) => setPw(e.target.value)}
                             type="password" name={"password"} autoFocus placeholder={"パスワード"}/>
                </div>
                <div className="flex justify-center my-2">
                    <button
                        className="disabled : opacity-40 py-2 px-4 rounded text-white bg-ingigo-600 "
                        desabled={!email || !pw}
                        type={"submit"}
                    >
                        {isLogin ? "ログイン" : "新規登録"}
                    </button>
                </div>
            </form>
        </div>
    );
};

export default Auth;