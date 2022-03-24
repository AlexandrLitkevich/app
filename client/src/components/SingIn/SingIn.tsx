import {SingInWrapper} from "./SingIn.styles";

import {Button, Form, Input, Typography} from 'antd';
import { useNavigate } from "react-router-dom";
import axios from "axios";

import { endPoint } from "../../constants"
import { FC } from "react";


const { Title } = Typography;

type Props = {
    setStatus: ( arg: boolean ) => void;
};

export const SingIn:FC<Props> = ({ setStatus }) => {
    // костыль роутинга


    const onFinish = (values: any) => {        
        axios.post(`${endPoint}/api/auth/`, values).then((res) => {
             if((res.status === 200)) {
                //TODO нехорошо
                setStatus(true)
             }
        })      
    };

    const onFinishFailed = (errorInfo: any) => {
        console.log('Failed:', errorInfo);
    };

    return (
        <SingInWrapper>
            <Title>Авторизация</Title>
            <Form
                name="basic"
                labelCol={{ span: 8 }}
                wrapperCol={{ span: 16 }}
                initialValues={{ remember: true }}
                onFinish={onFinish}
                onFinishFailed={onFinishFailed}
                autoComplete="off"
            >
                <Form.Item
                    label="Username"
                    name="username"
                    rules={[{ required: true, message: 'Please input your username!' }]}
                >
                    <Input />
                </Form.Item>

                <Form.Item
                    label="Password"
                    name="password"
                    rules={[{ required: true, message: 'Please input your password!' }]}
                >
                    <Input.Password />
                </Form.Item>

                <Form.Item wrapperCol={{ offset: 8, span: 16 }}>
                    <Button type="primary" htmlType="submit">
                        Sign In
                    </Button>
                </Form.Item>
            </Form>

        </SingInWrapper>
    )
}