import { FC } from "react"
import {Col, Row} from "antd";
import { ImageBlock } from "./Auth.styles";
import { SingIn } from "../SingIn";


type Props = {
    setStatus: (arg:boolean) => void;
}

export const Auth:FC<Props> = ({ setStatus }) => {

    return (
        <Row align="middle" >
            <Col span={16}>
                <div>
                    <ImageBlock/>
                </div>
            </Col>
            <Col span={8}>
                <Row justify="center">
                    <Col>
                        <SingIn setStatus={ setStatus }/>
                    </Col>
                </Row>
            </Col>
        </Row>
    )
}