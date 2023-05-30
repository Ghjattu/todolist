import React from 'react';
import PropTypes from 'prop-types';
import { useField } from 'formik';
import './TextInput.css';

const TextInput = ({ label, ...props }) => {
	const [field, meta] = useField(props);

	return (
		<div className="input-container">
			<label className="input-label" htmlFor={ props.id || props.name }>{ label }</label>
			<input className="input-field" { ...field } { ...props }/>
			{ meta.touched && meta.error &&
				<p className="input-error">{ meta.error }</p>
			}
		</div>
	);
};

TextInput.propTypes = {
	label: PropTypes.string.isRequired,
	id: PropTypes.string,
	name: PropTypes.string,
};

export default TextInput;
