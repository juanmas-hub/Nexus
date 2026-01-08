import { useState } from 'react';
import { api } from '../api/axios';

interface RegisterFormProps {
    onSwitchLogin: () => void;
}

const RegisterForm = ({ onSwitchLogin }: RegisterFormProps) => {
    const [formData, setFormData] = useState({
        email: '',
        password: '',
        first_name: '',
        last_name: ''
    });
    const [error, setError] = useState('');
    const [success, setSuccess] = useState(false);
    const [isLoading, setIsLoading] = useState(false);

    const handleSubmit = async (e: React.FormEvent) => {
        e.preventDefault();
        setError('');
        setIsLoading(true);

        try {
            await api.post('/register', formData);
            setSuccess(true);
            setTimeout(() => onSwitchLogin(), 2000);
        } catch (err: any) {
            setError(err.response?.data?.error || 'Error al registrar usuario');
        } finally {
            setIsLoading(false);
        }
    };

    return (
        <div className="flex min-h-screen items-center justify-center bg-gray-900 px-4">
            <div className="w-full max-w-md space-y-8 rounded-xl bg-gray-800 p-10 shadow-2xl border border-gray-700">
                <div className="text-center">
                    <h2 className="mt-6 text-3xl font-bold tracking-tight text-white">Crear Cuenta</h2>
                    <p className="mt-2 text-sm text-gray-400">Únete a la red de eventos Nexus</p>
                </div>

                {success ? (
                    <div className="rounded-md bg-green-500/10 p-4 text-sm text-green-500 text-center border border-green-500/20">
                        ¡Cuenta creada! Redirigiendo al login...
                    </div>
                ) : (
                    <form className="mt-8 space-y-4" onSubmit={handleSubmit}>
                        {error && (
                            <div className="rounded-md bg-red-500/10 p-3 text-sm text-red-500 text-center border border-red-500/20">
                                {error}
                            </div>
                        )}
                        <div className="grid grid-cols-2 gap-4">
                            <input
                                type="text"
                                placeholder="Nombre"
                                required
                                value={formData.first_name}
                                onChange={(e) => setFormData({...formData, first_name: e.target.value})}
                                className="block w-full rounded-md border-0 bg-gray-700 py-3 px-3 text-white ring-1 ring-inset ring-gray-600 focus:ring-2 focus:ring-indigo-500 sm:text-sm"
                            />
                            <input
                                type="text"
                                placeholder="Apellido"
                                required
                                value={formData.last_name}
                                onChange={(e) => setFormData({...formData, last_name: e.target.value})}
                                className="block w-full rounded-md border-0 bg-gray-700 py-3 px-3 text-white ring-1 ring-inset ring-gray-600 focus:ring-2 focus:ring-indigo-500 sm:text-sm"
                            />
                        </div>
                        <input
                            type="email"
                            placeholder="Email"
                            required
                            value={formData.email}
                            onChange={(e) => setFormData({...formData, email: e.target.value})}
                            className="block w-full rounded-md border-0 bg-gray-700 py-3 px-3 text-white ring-1 ring-inset ring-gray-600 focus:ring-2 focus:ring-indigo-500 sm:text-sm"
                        />
                        <input
                            type="password"
                            placeholder="Contraseña"
                            required
                            value={formData.password}
                            onChange={(e) => setFormData({...formData, password: e.target.value})}
                            className="block w-full rounded-md border-0 bg-gray-700 py-3 px-3 text-white ring-1 ring-inset ring-gray-600 focus:ring-2 focus:ring-indigo-500 sm:text-sm"
                        />

                        <button
                            type="submit"
                            disabled={isLoading}
                            className="w-full rounded-md bg-indigo-600 py-3 text-sm font-semibold text-white hover:bg-indigo-500 transition-colors disabled:opacity-50"
                        >
                            {isLoading ? 'Registrando...' : 'Registrarse'}
                        </button>

                        <p className="text-center text-sm text-gray-400">
                            ¿Ya tienes cuenta?{' '}
                            <button 
                                type="button"
                                onClick={onSwitchLogin}
                                className="font-semibold text-indigo-400 hover:text-indigo-300"
                            >
                                Inicia sesión
                            </button>
                        </p>
                    </form>
                )}
            </div>
        </div>
    );
};

export default RegisterForm;