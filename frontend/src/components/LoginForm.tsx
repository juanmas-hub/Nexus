import { useState } from 'react';
//import { api } from '../api/axios';
import { useAuth } from '../context/AuthContext';

interface LoginFormProps {
    onSwitchRegister: () => void;
}

const LoginForm = ({ onSwitchRegister }: LoginFormProps) => {
    const { login } = useAuth();
    const [email, setEmail] = useState('');
    const [password, setPassword] = useState('');
    const [error, setError] = useState('');
    const [isLoading, setIsLoading] = useState(false);

    const handleSubmit = async (e: React.FormEvent) => {
        e.preventDefault();
        setError('');
        setIsLoading(true);

        try {
            await login(email, password); 
        } catch (err: any) {
            setError(err.response?.data?.error || 'Credenciales inválidas');
        } finally {
            setIsLoading(false);
        }
    };

    return (
        <div className="flex min-h-screen items-center justify-center bg-gray-900 px-4">
            <div className="w-full max-w-md space-y-8 rounded-xl bg-gray-800 p-10 shadow-2xl border border-gray-700">
                <div className="text-center">
                    <h2 className="mt-6 text-3xl font-bold tracking-tight text-white">Nexus Events</h2>
                    <p className="mt-2 text-sm text-gray-400">Inicia sesión para continuar</p>
                </div>

                <form className="mt-8 space-y-6" onSubmit={handleSubmit}>
                    {error && (
                        <div className="rounded-md bg-red-500/10 p-3 text-sm text-red-500 text-center border border-red-500/20">
                            {error}
                        </div>
                    )}
                    <div className="space-y-4">
                        <input
                            type="email"
                            required
                            value={email}
                            onChange={(e) => setEmail(e.target.value)}
                            className="block w-full rounded-md border-0 bg-gray-700 py-3 px-3 text-white ring-1 ring-inset ring-gray-600 placeholder:text-gray-400 focus:ring-2 focus:ring-indigo-500 sm:text-sm"
                            placeholder="Email"
                        />
                        <input
                            type="password"
                            required
                            value={password}
                            onChange={(e) => setPassword(e.target.value)}
                            className="block w-full rounded-md border-0 bg-gray-700 py-3 px-3 text-white ring-1 ring-inset ring-gray-600 placeholder:text-gray-400 focus:ring-2 focus:ring-indigo-500 sm:text-sm"
                            placeholder="Contraseña"
                        />
                    </div>

                    <button
                        type="submit"
                        disabled={isLoading}
                        className="w-full rounded-md bg-indigo-600 py-3 text-sm font-semibold text-white hover:bg-indigo-500 transition-colors disabled:opacity-50"
                    >
                        {isLoading ? 'Cargando...' : 'Entrar'}
                    </button>

                    <p className="text-center text-sm text-gray-400">
                        ¿No tienes cuenta?{' '}
                        <button 
                            type="button"
                            onClick={onSwitchRegister}
                            className="font-semibold text-indigo-400 hover:text-indigo-300"
                        >
                            Regístrate aquí
                        </button>
                    </p>
                </form>
            </div>
        </div>
    );
};

export default LoginForm;