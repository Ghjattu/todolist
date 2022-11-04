import React from 'react';
import PropTypes from 'prop-types';
import './Button.css';

const Button = ({ type, text, ...props }) => {
	return (
		<button type={ type } { ...props }>{ text }</button>
	);
};

Button.propTypes = {
	type: PropTypes.string.isRequired,
	text: PropTypes.string.isRequired,
};

export default Button;
