import {useEffect, useState} from "react";
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

export const Main = () => {
    const [data, setData] = useState([])

    useEffect(() => {
        axios.get(`${endPoint}/users`).then(res =>{
            setData(res.data)
        })
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