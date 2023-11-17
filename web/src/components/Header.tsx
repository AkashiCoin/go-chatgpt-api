import {useContext} from "react";
import logo from "@/assets/react.svg";
import { Link, useNavigate } from 'react-router-dom';
import { Container, Dropdown, Icon, Menu, SemanticICONS } from 'semantic-ui-react';
import { UserContext} from "@/context/User";
import {getSystemName, showSuccess} from "@/helpers/utils.tsx";
import '../index.css';
import {UserActionTypes} from "@/types/user.ts";

// Header Buttons
const headerButtons: {name: string, to: string, icon: SemanticICONS}[] = [
    {
        name: '首页',
        to: '/',
        icon: "home" as SemanticICONS
    },
    {
        name: '设置',
        to: '/setting',
        icon: "setting" as SemanticICONS
    },
    {
        name: '关于',
        to: '/about',
        icon: "info circle" as SemanticICONS
    }
];



const Header = () => {
    const {state: userState , dispatch: userDispatch} = useContext(UserContext);
    const systemName = getSystemName();
    const navigate = useNavigate();

    async function logout() {
        // await API.get('/api/user/logout');
        showSuccess('注销成功!');
        userDispatch({
            type: UserActionTypes.LOGIN,
            payload: undefined
        });
        localStorage.removeItem('user');
        navigate('/login');
    }

    const renderButtons = () => {
        return headerButtons.map((button) => {
            return (
                <Menu.Item key={button.name} as={Link} to={button.to}>
                    <Icon name={button.icon} />
                    {button.name}
                </Menu.Item>
            );
        });
    };
    return (
        <>
            <Menu borderless style={{ borderTop: 'none' }}>
                <Container>
                    <Menu.Item as={Link} to='/' className={'hide-on-mobile'}>
                        <img src={logo} alt='logo' style={{ marginRight: '0.75em' }} />
                        <div style={{ fontSize: '20px' }}>
                            <b>{systemName}</b>
                        </div>
                    </Menu.Item>
                    {renderButtons()}
                    <Menu.Menu position='right'>
                        {userState.user ? (
                            <Dropdown
                                text={userState.user.username}
                                pointing
                                className='link item'
                            >
                                <Dropdown.Menu>
                                    <Dropdown.Item onClick={logout}>注销</Dropdown.Item>
                                </Dropdown.Menu>
                            </Dropdown>
                        ) : (
                            <Menu.Item
                                name='登录'
                                as={Link}
                                to='/login'
                                className='btn btn-link'
                            />
                        )}
                    </Menu.Menu>
                </Container>
            </Menu>
        </>
    );
}

export default Header