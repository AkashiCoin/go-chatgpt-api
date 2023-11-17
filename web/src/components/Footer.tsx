import { useEffect, useState } from 'react';

import { Container, Segment } from 'semantic-ui-react';
import { getFooterHTML, getSystemName } from '@/helpers/utils.tsx';

const Footer = () => {
    const systemName = getSystemName();
    const [footer, setFooter] = useState(getFooterHTML());
    let remainCheckTimes = 5;

    const loadFooter = () => {
        const footer_html = localStorage.getItem('footer_html');
        if (footer_html) {
            setFooter(footer_html);
        }
    };

    useEffect(() => {
        const timer = setInterval(() => {
            if (remainCheckTimes <= 0) {
                clearInterval(timer);
                return;
            }
            remainCheckTimes--;
            loadFooter();
        }, 200);
        return () => clearTimeout(timer);
    }, []);

    return (
        <Segment vertical>
            <Container textAlign='center'>
                {footer ? (
                    <div
                        className='custom-footer'
                        dangerouslySetInnerHTML={{ __html: footer }}
                    ></div>
                ) : (
                    <div className='custom-footer'>
                        <a
                            href='https://github.com/AkashiCoin/gin-template'
                            target='_blank'
                        >
                            {systemName} {import.meta.env.VITE_VERSION}{' '}
                        </a>
                    </div>
                )}
            </Container>
        </Segment>
    );
};

export default Footer;
