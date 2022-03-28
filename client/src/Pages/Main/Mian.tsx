import {FC, useEffect, useState} from "react";
import axios from "axios";
import { endPoint } from "../../constants";
import { Layout } from "antd";
import { AddedUrlForm } from "../../components/Form";
import { ListUrl } from "../../components/ListUrl";


export type DataType = {
    key: number
    username: string
    url: string

};

export const Main:FC = () => {
    const [data, setData] = useState([])
    const [userInfo, setInfoUser] = useState([])
    console.log("userInfo",userInfo);
    
    //BAD
    const getUserInfo = async () => {
        let token = localStorage.getItem("token")
        if (token) {
            await axios.get(`${endPoint}/api/userInfo`, {
                headers: {
                    Authorization: token
                }
            }).then(res => {
                setInfoUser(res.data)
            });
            axios.get(`${endPoint}/users`).then(res => {
                setData(res.data)
            });
        }
    };
    

    useEffect(() => {
        getUserInfo()
    }, [])

    const { Header, Content } = Layout;
    return (
        <Layout>
            <Header>
                <AddedUrlForm setUsers={setData}/>
            </Header>
            <Content>
                <ListUrl setUsers={setData} users={data}/>
            </Content>
        </Layout>
    )


};